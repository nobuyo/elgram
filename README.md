# elgram

element-gram

## Usage

```
% elgram <source> <indices>
```

```sh
% elgram 5325733 3212
# => 2353

% elgram somephrase 4321
# => emos

# You'd like to use over 10 characters source, give comma splitted indices.
% elgram super-long-phrase 12,5,4,4
# => pree
```

## Build

```sh
% go build -o elgram main.go
```
