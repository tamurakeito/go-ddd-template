echo '\x1b[36mプロジェクト名を入力しEnterキーを押してください\x1b[0m'
read -p '> ' title

rm -r ./.init.sh
echo '#' ${title} > ./README.md
echo '...init操作関連ファイルをクリアしました...'

yes | rm -r ./.git
echo '...gitをリセットしました...'

# packageにある名前を一括で変更する記述
find . -type f -exec sed -i '' "s/go-ddd-template/${title}/g" {} +
cd ../
mv go-ddd-template ${title}
echo '...プロジェクト名を設定しました...'
