
build_mains(){
    for d in ./${1}/*/ ; do (cd "$d" && go build main.go); done
}

build_mains coffeemakingtoaster