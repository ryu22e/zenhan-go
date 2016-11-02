# zenhan-go
Go Library to convert between Zenkaku(fullwidth Japanese) and Hankaku(halfwidth Japanese)

This is clone of [zenhan-py](https://github.com/sspiral/zenhan-py)

## How to install

Run the following command.

    go get -u github.com/ryu22e/zenhan-go

## Usage

    package main

    import (
        "fmt"
        "github.com/ryu22e/zenhan-go"
    )

    const (
        hankaku = "ABC123ｱｲｳﾊﾟﾋﾟﾌﾟ"
        zenkaku = "ＡＢＣ１２３アイウ"
    )

    func main() {
        fmt.Println(zenhan.H2z(hankaku, zenhan.ASCII))
        // ＡＢＣ123ｱｲｳﾊﾟﾋﾟﾌﾟ
        fmt.Println(zenhan.H2z(hankaku, zenhan.DIGIT))
        // ABC１２３ｱｲｳﾊﾟﾋﾟﾌﾟ
        fmt.Println(zenhan.H2z(hankaku, zenhan.KANA))
        // ABC123アイウパピプ
        fmt.Println(zenhan.H2z(hankaku, zenhan.ALL))
        // ＡＢＣ１２３アイウパピプ
        fmt.Println(zenhan.H2z(hankaku, zenhan.ALL, "A", "1", "ｱ"))
        // AＢＣ1２３ｱイウパピプ

        fmt.Println(zenhan.Z2h(zenkaku, zenhan.ASCII))
        // ABC１２３アイウ
        fmt.Println(zenhan.Z2h(zenkaku, zenhan.DIGIT))
        // ＡＢＣ123アイウ
        fmt.Println(zenhan.Z2h(zenkaku, zenhan.KANA))
        // ＡＢＣ１２３ｱｲｳ
        fmt.Println(zenhan.Z2h(zenkaku, zenhan.ALL))
        // ABC123ｱｲｳ
        fmt.Println(zenhan.Z2h(zenkaku, zenhan.ALL, "Ａ", "１", "ア"))
        // ＡBC１23アｲｳ
    }

[![wercker status](https://app.wercker.com/status/11360f18dee678ff652cfb579e4fdd5e/s/master "wercker status")](https://app.wercker.com/project/byKey/11360f18dee678ff652cfb579e4fdd5e)
