package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // クエリパラメータを解析する
    query := r.URL.Query()
    name := query.Get("name")

    // レスポンス用のマップを作成
    response := map[string]string{
        "message": "Hello " + name,
    }

    // Content-Typeヘッダーをapplication/jsonに設定
    w.Header().Set("Content-Type", "application/json")

    // マップをJSONにエンコードしてレスポンスとして送信
    json.NewEncoder(w).Encode(response)
}

// カテゴリの配列を含むJSONを返す
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
    categories := []string{"Technology", "Science", "Sports", "Health", "Entertainment"}

    // レスポンス用のマップを作成
    response := map[string]interface{}{
        "categories": categories,
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(response)
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // クエリを調べる
    query := r.URL.Query()
    ope := query.Get("o")

    // x y が数字でなければエラー
    x, err := strconv.Atoi(query.Get("x"))
    if err != nil {
		// strは数字ではない
		fmt.Println("Not a number")
	}

    y, err := strconv.Atoi(query.Get("y"))
    if err != nil {
		// strは数字ではない
		fmt.Println("Not a number")
	}

    fmt.Println("ope " + ope)
    fmt.Println("x ")
    fmt.Println(x)
    fmt.Println("y ")
    fmt.Println(y)

    // 計算する
    var result int
    switch ope {
    case "+":
        result = x + y
    case "-":
        result = x - y
    case "*":
        result = x * y
    case "/":
        if y != 0 {
            result = x / y
        } else {
            response := map[string]interface{}{
                "Error": "Division by zero",
            }
            json.NewEncoder(w).Encode(response)
            return
        }
    default:
        json.NewEncoder(w).Encode("Error: Invalid operator")
        return
    }

    // 計算結果をjsonで返す
    response := map[string]interface{}{
        "answer": result,
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
    fmt.Println("Starting the server!")

    // ルートとハンドラ関数を定義
    http.HandleFunc("/api/hello", helloHandler)

    http.HandleFunc("/api/categories", categoriesHandler)

    http.HandleFunc("/api/calculator", calculatorHandler) // /api/calculator?o={operator}&x={x}&y={y}
    // 演算子はエンコードしたものを渡さないとスペースと判断される
    // + %2B, - %2D, * %2A, / %2F

    // エラー /api/calculator?o=+&x=3&y=12       "+"はURLエンコーディングではスペースと判断される

    // 足し算: api/calculator?o=%2B&x=3&y=12 -> {"answer":15}
    // 割り算: api/calculator?o=%2F&x=30&y=5 -> {"answer":6}
    // 0での割り算: api/calculator?o=%2F&x=3&y=0  -> {"Error":"Division by zero"}


    // 8000番ポートでサーバを開始
    http.ListenAndServe(":8000", nil)
}
