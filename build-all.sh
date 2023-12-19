
build_mains(){
    for d in ./${1}/*/ ; do (cd "$d" &&  echo "Building $d" && go build -v main.go); done
}

build_mains coffeemakingtoaster
build_mains LarsFlieger