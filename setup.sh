#!/bin/bash

# Prompt untuk nama module baru
echo "Masukkan nama module Anda (contoh: github.com/username/nama-project atau nama-project):"
read module_name

if [ -z "$module_name" ]; then
  echo "Nama module tidak boleh kosong. Coba lagi."
  exit 1
fi

# Ganti nama module di seluruh file .go
echo "Mengganti 'github.com/yudhisrana/boilerplate-goyam' dengan '$module_name' di semua file..."
find . -type f -name "*.go" -exec sed -i "s/github.com\/yudhisrana\/boilerplate-goyam/$module_name/g" {} +

# Inisialisasi module baru
echo "Inisialisasi go module dengan nama $module_name..."
go mod init "$module_name"
go mod tidy

echo "Setup selesai!"
