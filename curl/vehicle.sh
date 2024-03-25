#!/bin/bash

# Criar um novo veículo
curl -X POST http://localhost:8080/vehicles \
-H 'Content-Type: application/json' \
-d '{
    "name": "Toyota Corolla",
    "model": "2020",
    "status": "Novo",
    "color": "Azul",
    "mileage": 0,
    "bodyType": "Sedan",
    "transmission": "Automático",
    "fuelType": "Gasolina",
    "doors": 4,
    "review": true,
    "price": 20000.00,
    "description": "Um carro novo e confortável"
}'

echo "Veículo criado"

# Substitua {id} pelo ID real do veículo nas próximas chamadas
# Buscar um veículo específico
curl -X GET http://localhost:8080/vehicles/{id}
echo "Veículo consultado"

# Listar todos os veículos
curl -X GET http://localhost:8080/vehicles
echo "Lista de veículos recuperada"

# Atualizar um veículo existente
curl -X PUT http://localhost:8080/vehicles/{id} \
-H 'Content-Type: application/json' \
-d '{
    "name": "Toyota Corolla",
    "model": "2021",
    "status": "Usado",
    "color": "Preto",
    "mileage": 5000,
    "bodyType": "Sedan",
    "transmission": "Manual",
    "fuelType": "Gasolina",
    "doors": 4,
    "review": false,
    "price": 19000.00,
    "description": "Em ótimas condições, com poucos quilômetros rodados"
}'
echo "Veículo atualizado"

# Deletar um veículo
curl -X DELETE http://localhost:8080/vehicles/{id}
echo "Veículo deletado"
