$ time go run teiler.go > 100000.txt
real    0m23.850s
user    0m23.492s
sys     0m0.391s


$ time go run go-teiler.go > 100000-got.txt
real    0m12.118s
user    0m40.751s
sys     0m0.667s

$ diff 100000.txt 100000-got.txt
$
