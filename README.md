# Stress Testing Tool

## Abordagem

O package ```stresstest``` contém a lógica principal da aplicação.
Utilizando de golang channels para controlar a concorrecia, disparamos
goroutines em paralelo para atingir a concorrencia desejada.

Para controlar que todas as requisições foram executadas, utilizo waitgroups.

## Usage

Para utilizar a ferramenta, é necessário passar três parâmetros: url, requests e concurrency.

Aqui está um exemplo de como você pode executar a ferramenta:

```bash
docker build . -t "alusoft-stress-test" && \
docker run alusoft-stress-test --url=http://example.com --requests=10 --concurrency=5
```

Neste exemplo, a ferramenta realizará um total de 10 requisições para http://example.com com um nível de concorrência de 5.

## Output

Após executar o teste, você verá uma saída como esta:

```
Stress Testing http://example.com, Requests 10, Concurrency 5...
Total of request with Status code[200]: 10 
Total of request with: 10 
Execution time 1.056415s
```

Isso significa que a ferramenta fez um total de 1000 requisições para http://example.com e todas essas requisições tiveram um código de status HTTP 200 (OK). O tempo de execução é o tempo que levou para o teste ser concluído.