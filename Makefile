# Название исполняемого файла для Windows
#BINARY=secLab.exe

#Название исполняемого файла для Linux
BINARY=secLab

# Команда для сборки
build:
	go build -o $(BINARY) .

# Команда для запуска
run: build
	$(BINARY)

# Команда для тестирования
test:
	go test ./...

# Команда для очистки (удаления скомпилированного файла)
clean:
	del $(BINARY)

# По умолчанию выполнить команду run
.PHONY: build run test clean
