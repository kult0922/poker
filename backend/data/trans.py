import re

# 入力ファイルを読み込む
with open('input.txt', 'r') as f:
    text = f.read()

# 2つ以上のスペースを,に置換する
text = re.sub(' {2,}', ',', text)

# 出力ファイルに書き込む
with open('output.txt', 'w') as f:
    f.write(text)
