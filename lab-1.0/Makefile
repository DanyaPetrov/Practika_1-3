# Makefile для сборки программы и создания deb-пакета

.PHONY: build clean deb

# Путь к Go-модулю
MODULE := lab

# Сборка бинарного файла (если нужно вручную)
build:
	go build -o $(MODULE) ./cmd/$(MODULE)

test:
	@echo "Запуск тестов..."
	go test -v ./...

# Сборка deb-пакета с использованием официальной инфраструктуры Debian
deb:
	dpkg-buildpackage -us -uc -b

# Очистка временных файлов
clean:
	rm -f $(MODULE)
	dpkg-buildpackage -T clean 2>/dev/null || true
