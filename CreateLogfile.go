// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package srapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/term"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。
Vwe.2.0.0 引数（ログファイルのプリフィックス）を可変長とする。戻り値にerrを追加する。

*/

// ログファイルを作る。
func CreateLogfile(dsc ...string) (logfile *os.File, err error) {

	// 1. 現在の作業ディレクトリのパスを取得
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("現在の作業ディレクトリの取得に失敗しました: %v", err)
	}
	// 2. パスのベース名（最後のディレクトリ名）を抽出
	//    filepath.Base は、パスの最後の要素（ファイル名またはディレクトリ名）を返します。
	//    例: "/home/user/myproject" -> "myproject"
	//    例: "/home/user/myproject/" -> "myproject" (末尾のスラッシュは無視される)
	baseName := filepath.Base(currentDir)

	//      ログファイルの設定
	// logfilename := os.Args[0]
	logfilename := baseName
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

	// --------------------------------

	// フォアグラウンド（端末に接続されているか）を判定
	isForeground := term.IsTerminal(int(os.Stdout.Fd()))
	if isForeground {
		// フォアグラウンドならログファイル + コンソール
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	} else {
		// バックグラウンドならログファイルのみ
		log.SetOutput(logfile)
	}

	log.SetFlags(log.Lmicroseconds)
	log.Printf("Version=%s Start\n", Version)

	return
}
