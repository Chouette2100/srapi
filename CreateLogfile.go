/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package srapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。
Vwe.2.0.0 引数（ログファイルのプリフィックス）を可変長とする。戻り値にerrを追加する。

*/

//	ログファイルを作る。
func CreateLogfile(dsc... string) (logfile *os.File, err error) {
	//      ログファイルの設定
	logfilename := os.Args[0]
	for _, dsci := range dsc {
		logfilename += "_" + dsci
	}
	logfilename += "_" + time.Now().Format("20060102")
	logfilename += ".txt"
	logfile, err = os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		err = fmt.Errorf("CreateLogfile(): %w", err)
		return
	}

	//      log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	return
}
