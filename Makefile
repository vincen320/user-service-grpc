# Build proto (only requires proto/user.proto to exist)
protoBuild: proto/user.proto
	protoc --proto_path=proto --grpc-gateway_out=proto/ \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=source_relative \
	--go_out=proto/ --go_opt=paths=source_relative \
	--go-grpc_out=proto/ --go-grpc_opt=paths=source_relative \
	user.proto

#jangan lupa sebelum '\' pakai spasi

#di VSCode tabnya masih terhitung spasi 4x, sedangkan Makefile harus beneran bentuk Tab, jadi cara ubahnya adalah pakai vim\
vim adalah editor yang pakai command editny, jadi semuany pakai command\
vim makefile\
lalu didalam vim ketik\
:%s/^[ ]\+/\t/g\
selalu diawali dengan ":" untuk commadnny\
lalu untuk save(w) dan exit(q) adalah\
:wq

#ringkasan\
vim makefile\
:%s/^[ ]\+/\t/g\
:wq\


#https://stackoverflow.com/questions/1789594/how-do-i-write-the-cd-command-in-a-makefile
#https://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date || jangan buat nama file sama dengan namatarget di make file

#Masalah import failed di user.proto bisa diatasi dengan mendefinisikan --proto_path=proto, jadi protoc akan cari file protocnya dimulai dengan path sekarang(path Makefile) + path tersebut.
#jadi waktu user.proto langsung saja, tidak perlu proto/user.proto karena --protocnya sudah masuk ke dalam folder proto, tetapi untuk definisi go_out nya sesuai penjelasan(#vinPath)
#https://stackoverflow.com/questions/64048132/proto-path-passed-empty-directory-name || -I . dan -proto_path=. adalah sama
#You have to use the --proto_path command-line flag (aka -I) to tell protoc where to look for .proto files. If you don't provide a path, by default it will only search the current directory.

#vinPath
#karena makefile diluar, jadi kalau path=. , jadinya sesuai dengan posisi Makefile, jadi definisikan pathnya sesuai dengan posisi Makefile.