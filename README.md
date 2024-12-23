# ЭФМО-02-24 Шатов Д.Я. Дергачев А.С. Итоговый проект
Тема: онлайн магазин/библиотека компьютерных игр (на вроде Steam)

Сборка докера

docker-compose up --build

Генерация для swagger

swag init -g ./cmd/app/main.go -o api/docs