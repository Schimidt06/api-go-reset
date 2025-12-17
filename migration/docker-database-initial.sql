-- Criação da Tabela 'personalidades'
-- Essa tabela vai mapear diretamente com a Struct que criaremos no Go.
create table personalidades(
    -- id serial primary key:
    -- 1. 'serial': No Postgres, isso cria um autoincremento (1, 2, 3...).
    --    No Go, você vai ignorar esse campo na hora de criar (INSERT),
    --    pois o banco gera sozinho.
    -- 2. 'primary key': Garante que nunca existirão dois IDs iguais.
    --    É o índice principal para buscas rápidas (.First(&p, id)).
    id serial primary key,

    -- nome varchar:
    -- 'varchar' sem tamanho definido no Postgres aceita qualquer tamanho de texto.
    -- No Go, isso será mapeado para o tipo 'string'.
    nome varchar,

    -- historia varchar:
    -- Mesma coisa, guardará o texto longo da biografia.
    historia varchar
);

-- Inserção de Dados (Seed):
-- Isso é o que chamamos de 'Seed' (Semente) ou carga inicial.
-- Serve para você não começar com o sistema vazio.
INSERT INTO personalidades(nome, historia) VALUES
(
    'Deodato Petit Wertheimer',
    'Deodato Petit Wertheimer foi um médico e político brasileiro...'
),
(
    'Carmela Dutra',
    'Carmela Teles Leite Dutra foi a primeira-dama do Brasil...'
);