#!/bin/bash

modules=(
  "app/checkout"
  "app/payment"
  "app/auth"
  "app/cart"
  "app/frontend"
  "app/order"
  "app/product"
  "app/user"
)

for module in "${modules[@]}"; do
  (
    echo "🛠️  Building $module"
    cd "$module" && go build "github.com/All-Done-Right/douyin-mall-microservice/$module"
    cd ..
  ) &
done

wait
echo "✅ All modules built successfully"
