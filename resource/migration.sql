CREATE TABLE pokemon (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    tipo VARCHAR(50) NOT NULL,
    estagio INT,
    habilidade VARCHAR(255),
    ataque VARCHAR(255) not null,
    ps INT NOT NULL,
    recuo INT,
    fraqueza VARCHAR(50),
    CONSTRAINT fk_habilidade FOREIGN KEY (habilidade) REFERENCES habilidade(nome_habilidade),
    CONSTRAINT fk_ataque FOREIGN KEY (ataque) REFERENCES ataque(nome_ataque)
);

CREATE TABLE ataque (
    nome_ataque VARCHAR(255) PRIMARY KEY,
    dano_ataque INT NOT NULL,
    custo_ataque VARCHAR(50),
    efeito_ataque TEXT
);

CREATE TABLE habilidade (
    nome_habilidade VARCHAR(255) PRIMARY KEY,
    efeito_habilidade TEXT
);


CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    efeito TEXT
);

CREATE TABLE apoiador (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    efeito TEXT
);