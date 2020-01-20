CREATE TABLE contato (
    id serial,
    nome VARCHAR(40) NOT NULL,
    numero VARCHAR(11) NOT NULL,
    random INT,
    PRIMARY KEY(id)
);

INSERT INTO contato (nome, numero) VALUES
('Vanessa da mata', '19992345456'),
('Negretude Júnior', '19992345678'),
('Jack Bauer', '19998785675');