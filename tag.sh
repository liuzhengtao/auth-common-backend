na=$1
info=$2

git tag -d $na
git push origin :refs/tags/$na
git tag -a $na -m $info
git push origin $na