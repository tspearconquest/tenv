/*
 *
 * Copyright 2024 tofuutils authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package flatparser

import (
	"bytes"
	"errors"
	"io/fs"
	"os"

	"github.com/tofuutils/tenv/v4/config"
	"github.com/tofuutils/tenv/v4/pkg/loghelper"
	"github.com/tofuutils/tenv/v4/versionmanager/semantic/types"
)

func NoMsg(_ loghelper.Displayer, value string, _ string) string {
	return value
}

func Retrieve(filePath string, conf *config.Config, displayMsg func(loghelper.Displayer, string, string) string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		conf.Displayer.Log(loghelper.LevelWarnOrDebug(errors.Is(err, fs.ErrNotExist)), "Failed to read file", loghelper.Error, err)

		return "", nil
	}

	resolvedVersion := string(bytes.TrimSpace(data))
	if resolvedVersion == "" {
		return "", nil
	}

	return displayMsg(conf.Displayer, resolvedVersion, filePath), nil
}

func RetrieveVersion(filePath string, conf *config.Config) (string, error) {
	return Retrieve(filePath, conf, types.DisplayDetectionInfo)
}
