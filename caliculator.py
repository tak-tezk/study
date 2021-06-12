import tkinter as tk

# ボタンの配置を定義
BUTTONS = [
    ['(', ')', 'C', 'B'],
    ['7', '8', '9', '/'],
    ['4', '5', '6', '*'],
    ['1', '2', '3', '-'],
    ['0', '.', '=', '+']
]

# ボタンイベントの作成関数
def make_click(ch):
    def click(e):
        st = disp.get()
        length = len(st)
        if ch == '=': calc(0); return
        elif ch == 'B': disp.delete(length-1)
        elif ch == 'C': disp.delete(0, length)
        else: disp.insert(tk.END, ch)
    return click

# 計算式を計算
def calc(e):
    label["text"] = '= ' + str(eval(disp.get()))

# ウィンドウを作成
win = tk.Tk()
win.title("電卓")
win.geometry("320x430")

# ディスプレイ部分
disp = tk.Entry(win, font=('', 20), justify="center")
disp.pack(fill='x')
disp.bind('<Return>', calc)
label = tk.Label(win, font=('', 20), anchor="center")
label.pack(fill='x')

# 電卓のボタンを一括作成
fr = tk.Frame(win)
fr.pack()
for y, cols in enumerate(BUTTONS):
    for x, n in enumerate(cols):
        btn = tk.Button(fr, text=n,
            font=('', 20), width=4, height=2)
        btn.grid(row=y+1, column=x+1)
        btn.bind('<1>', make_click(n))

win.mainloop()
