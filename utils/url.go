package utils

import (
	"net/url"
	"regexp"

	"github.com/pkg/errors"
)

func UnescapeURL(rawURL string) string {
	if u, err := url.QueryUnescape(rawURL); err == nil {
		return u
	}
	return rawURL
}

func ValidateDocumentURL(url string) (string, string, error) {
	reg := regexp.MustCompile("^https://[\\w-.]+/(docs|docx|wiki)/([a-zA-Z0-9]+)")
	matchResult := reg.FindStringSubmatch(url)
	if matchResult == nil || len(matchResult) != 3 {
		return "", "", errors.Errorf("Invalid feishu/larksuite document URL pattern")
	}
	docType := matchResult[1]
	docToken := matchResult[2]
	return docType, docToken, nil
}

func ValidateFolderURL(url string) (string, error) {
	reg := regexp.MustCompile("^https://[\\w-.]+/drive/folder/([a-zA-Z0-9]+)")
	matchResult := reg.FindStringSubmatch(url)
	if matchResult == nil || len(matchResult) != 2 {
		return "", errors.Errorf("Invalid feishu/larksuite folder URL pattern")
	}
	folderToken := matchResult[1]
	return folderToken, nil
}

func ValidateWikiURL(url string) (string, string, error) {
	// 支持三种知识库URL格式：
	// 1. 知识库设置页面：https://xxx/wiki/settings/[token]
	// 2. 知识库空间页面：https://xxx/wiki/space/[spaceID]
	// 3. 知识库页面：https://xxx/wiki/[token]

	// 先尝试知识库设置页面格式
	settingsReg := regexp.MustCompile(`^(https://[\w-.]+)/wiki/settings/([a-zA-Z0-9]+)`)
	matchResult := settingsReg.FindStringSubmatch(url)
	if matchResult != nil && len(matchResult) == 3 {
		prefixURL := matchResult[1]
		wikiToken := matchResult[2]
		return prefixURL, wikiToken, nil
	}

	// 再尝试知识库空间页面格式
	spaceReg := regexp.MustCompile(`^(https://[\w-.]+)/wiki/space/([a-zA-Z0-9]+)`)
	matchResult = spaceReg.FindStringSubmatch(url)
	if matchResult != nil && len(matchResult) == 3 {
		prefixURL := matchResult[1]
		spaceID := matchResult[2]
		return prefixURL, spaceID, nil
	}

	// 最后尝试知识库页面格式
	pageReg := regexp.MustCompile(`^(https://[\w-.]+)/wiki/([a-zA-Z0-9]+)`)
	matchResult = pageReg.FindStringSubmatch(url)
	if matchResult != nil && len(matchResult) == 3 {
		prefixURL := matchResult[1]
		wikiToken := matchResult[2]
		return prefixURL, wikiToken, nil
	}

	return "", "", errors.Errorf("Invalid feishu/larksuite wiki URL pattern. Expected format: https://xxx/wiki/[token] or https://xxx/wiki/settings/[token] or https://xxx/wiki/space/[spaceID]")
}
