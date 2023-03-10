# Programowanie Full-Stack w Chmurze Obliczeniowej - Zadanie 1

1.  
Przygotowana aplikacja napisana została w jęzku Go i ma formę aplikacji internetowej.

Aby uruchomić aplikację lokalnie należy (po uprzednim pobraniu i zainstalowaniu Go) wykonać polecenie

```go
go run main.go
```

Do utworzenia repozytorium i przesłania do niego kodu wykorzystane zostały polecenia:

```
git init
git add .
git commit -m 'init'
gh repo create pfswcho-zad1 --public --source . --remote=pfswcho-zad1 --push
```
Utworzone repozytorium:
![Repo screenshot](https://user-images.githubusercontent.com/67523413/210115764-19854871-5b7e-4171-8401-11dace248b39.png)

2.
Do zbudowania obrazu lokalnie, wykorzystany został deamon Buildkit uruchomiony jako kontener oraz przygotowany [Dockerfile]().
Zbudowany przez BuildKit obraz wyeksportowany został jako Docker tarball oraz załadowany jako obraz poprzez polecenie `docker load`

Do uruchomienia kontenera z silnikiem BuildKit i zbudowania obrazu użyte zostały polecenia:
```
docker run --rm --privileged -d --name buildkit moby/buildkit
export BUILDKIT_HOST=docker-container://buildkit
buildctl build --frontend=dockerfile.v0 --local context=. --local dockerfile=. --output type=docker,name=zad1-test | docker load
```

Uruchomienie kontenera na podstawie obrazu:
```
docker run --rm -it -p 8080:8080 zad1-test
```
![Building image and running container](https://user-images.githubusercontent.com/67523413/210115193-65dad646-7a1c-465f-8364-d53f2713f096.png)

Po wejściu na adres [localhost:8080](http://localhost:8080) można przetestować działanie aplikacji:

![Working app](https://user-images.githubusercontent.com/67523413/210115191-f1f0cae7-b7de-496b-a20a-a2f084c6eee2.png)
![Working app with result](https://user-images.githubusercontent.com/67523413/210115190-a376b9b1-2529-4ed3-a428-58e105dbe2d0.png)

3.
Do wykonania tej części zadania wykorzystany został przygotowany plik definiujący workflow [fib.yml](https://github.com/xlpnsk/pfswcho-zad1/blob/main/.github/workflows/fib.yml). Kolejne zadania opisane są za pomocą komentarzy w pliku. Workflow uruchamiany jest poprzez event workflow_dispatch i wymaga podania parametru, który wykorzystywany jest do określenia wersji dla budowanego obrazu.

Na potrzeby workflow dodane zostały także zmienne secret w ustawieniach repozytorium:

![Secrets](https://user-images.githubusercontent.com/67523413/210115189-62932d51-bfd0-410b-a4f3-00c2b2d2fb8c.png)

4.
Do uruchomienia workflow za pomocą `gh` wykorzystane zostało polecenie `gh workflow run` z flagą `-f` pozwalającą na przekazanie wartości parametru. Do wylistowania aktywnych workflow oraz śledzenia uruchomionego wykorzystane zostały polecenia `gh run list` i `gh run watch`. Podsumowując, do wykonania niezbędnych działań wykonane zostały polecenia:
```
gh workflow run fib.yml -f version=1.0.0
gh run list
gh run watch <workflow_id>
```
![Successful workflow](https://user-images.githubusercontent.com/67523413/210115188-f8f214f6-78e8-40ae-930c-848fd79239ec.png)

Po udanym wykonaniu workflow można sprawdzić, że zarówno obraz kontenera (dla obydwu architektur), jak i cache zostały zapisane kolejno na repozytorium [ghcr.io (obraz)](https://github.com/xlpnsk/pfswcho-zad1/pkgs/container/pfswcho-zad1) oraz [docker.io (cache)](https://hub.docker.com/repository/docker/xlpnsk/pfswcho-zad1).

Można teraz sprawdzić poprawność zbudowanego obrazu poprzez wykonanie poleceń:
```
docker run --rm -it -p 8080:8080 ghcr.io/xlpnsk/pfswcho-zad1:1.0.0
```
![Running app](https://user-images.githubusercontent.com/67523413/210115186-1aa1edd6-ad25-4bc3-a6ff-3a9e31d78c49.png)
![Working app](https://user-images.githubusercontent.com/67523413/210115183-6e68df9b-2c1b-4cbf-95e5-e6d1cfe5594a.png)

Aby potwierdzić, że podczas budowania obrazu cache jest również importowany ze zdalnego registry docker.io, można uruchomić ponownie worflow (tym razem z poziomu przeglądarki na stronie platformy Github) i sprawdzić logi podczas wykonywania kolejnych zadań:

![Importing cache](https://user-images.githubusercontent.com/67523413/210115182-47a96bd3-2740-497f-994d-42e864bf7de1.png)

## Zadanie nieobowiązkowe 1
Na potrzeby zadania przygotowany został nowy plik Dockerfile [Dockerfile_dod1](https://github.com/xlpnsk/pfswcho-zad1/blob/main/Dockerfile_dod1) oraz plik workflow [fib_dod1.yml](https://github.com/xlpnsk/pfswcho-zad1/blob/main/.github/workflows/fib_dod1.yml).
Do pliku Dockerfile dodane zostały odpowiednie deklaracje, aby uwzględnić specyfikę budowy obrazu dla różnych platform sprzętowych, które zostały opisane w komentarzach do pliku. Plik workflow zmienił się nieznacznie - zmienione zostały generowane tagi, aby odróżnić budowane obrazy od tych z części pierwszej zadania oraz wskazany został nowy plik Dockerfile do budowy obrazów.

Po pozytywnym wykonaniu zadań nowego workflow oraz pobraniu i uruchomieniu obrazu można zobaczyć, że aplikacja działa poprawnie:

![Successful workflow after changes](https://user-images.githubusercontent.com/67523413/210115180-67cb1742-a80c-4c48-97ce-69614afbee3b.png)
![Running app after changes](https://user-images.githubusercontent.com/67523413/210115179-5b4d384b-a8e1-4f85-8ff8-243ad2fa5c32.png)
![Working app after changes](https://user-images.githubusercontent.com/67523413/210115178-fbd83190-2e1e-4dea-890a-0f259242b245.png)
