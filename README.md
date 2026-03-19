# PasswordEncoder

コマンドライン引数で受け取った文字列を bcrypt でハッシュ化して標準出力に表示します。

## 実行

```bash
./password-encoder my-password
```

## 補足

bcrypt はソルトを含むため、同じ入力でも毎回異なるハッシュ文字列が生成されます。  
ただし `bcrypt.CompareHashAndPassword` を使えば照合できます。

