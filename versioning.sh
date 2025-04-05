if [[ $1 != v* ]]; then
    v="v$1"
else
    v="$1"
fi

git tag $v
git push origin $v
cd framework
GOPROXY=proxy.golang.org go list -m github.com/truc-engine/truc@$v