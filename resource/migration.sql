CREATE TABLE IF NOT EXISTS pokemon (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    tipo VARCHAR(50) NOT NULL,
    estagio INT,
    geracao INT,
    ps INT NOT NULL,
    recuo INT,
    fraqueza VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS ataque (
    nome_ataque VARCHAR(255) PRIMARY KEY,
    dano_ataque INT NOT NULL,
    custo_ataque TEXT[],
    efeito_ataque TEXT
);

CREATE TABLE IF NOT EXISTS pokemon_ataque (
    id_pokemon INT,
    ataque VARCHAR (255),
    PRIMARY KEY (id_pokemon, ataque),
    FOREIGN KEY (id_pokemon) REFERENCES pokemon(id),
    FOREIGN KEY (ataque) REFERENCES ataque (nome_ataque)
);

CREATE TABLE IF NOT EXISTS habilidade (
    nome_habilidade VARCHAR(255) PRIMARY KEY,
    efeito_habilidade TEXT
);

CREATE TABLE IF NOT EXISTS pokemon_habilidade (
    id_pokemon INT,
    habilidade VARCHAR (255),
    PRIMARY KEY (id_pokemon, habilidade),
    FOREIGN KEY (id_pokemon) REFERENCES  pokemon(id),
    FOREIGN KEY (habilidade) REFERENCES habilidade(nome_habilidade)
);

CREATE TABLE IF NOT EXISTS item (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    efeito TEXT
);

CREATE TABLE IF NOT EXISTS apoiador (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    card_type VARCHAR(50) NOT NULL,
    efeito TEXT
);