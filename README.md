# Como usar esse Script

Crie um arquivo `secret.yaml` com as informações apenas do secret.

Note que deve conter apenas do data até o final dos segredos.
````yaml
data:
username: dXNlcm5hbWU=
password: cGFzc3dvcmQ=
````


Execute o código:

```bash
go run main.go secret.yaml
```

Ele vai gerar um arquivo output.txt em chave valor e com as chaves decriptadas.

A saída do arquivo deve ser assim:
```plaintext
username=username
password=password
```
