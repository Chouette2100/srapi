/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。
Ver.1.1.0 yaml.Unmarshal() を yaml.UnmarshalStrict() に変更する。

*/

// 設定ファイルを読み込む
//	以下の記事を参考にさせていただきました。
//		【Go初学】設定ファイル、環境変数から設定情報を取得する
//			https://note.com/artefactnote/n/n8c22d1ac4b86
//
func LoadConfig(filePath string, config interface{}) (err error) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("os.ReadFile: %w", err)
		return err
	}

	content = []byte(os.ExpandEnv(string(content)))
	//	log.Printf("content=%s\n", content)

	if err := yaml.UnmarshalStrict(content, config); err != nil {
		err = fmt.Errorf("yaml.Unmarshal(): %w", err)
		return err
	}

	//	log.Printf("\n")
	//	log.Printf("%+v\n", config)
	//	log.Printf("\n")

	return nil
}
