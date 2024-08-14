1. 名前空間の作成:
   ```
   kubectl apply -f manifests/namespace.yaml
   ```

2. ConfigMapとSecretの作成:
   ```
   kubectl apply -f manifests/configmaps/
   kubectl apply -f manifests/secrets/
   ```

3. ストレージの設定:
   ```
   kubectl apply -f manifests/storage/
   ```

4. PostgreSQLのデプロイ:
   ```
   kubectl apply -f manifests/deployments/postgres-deployment.yaml
   kubectl apply -f manifests/services/postgres-service.yaml
   ```

5. kakeibo-user-serviceのデプロイ:
   ```
   kubectl apply -f manifests/deployments/kakeibo-user-service-deployment.yaml
   kubectl apply -f manifests/services/kakeibo-user-service-service.yaml
   ```
