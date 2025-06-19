CREATE TABLE
    estabelecimento (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(50),
        endereco varchar(50),
        cnpj varchar(50),
        PRIMARY KEY (ID)
    );

CREATE TABLE
    consumidor (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(150),
        cpf varchar(50),
        PRIMARY KEY (ID)
    );

CREATE TABLE
    nota_fiscal (
        id integer NOT NULL AUTO_INCREMENT,
        data_emissao datetime,
        qtd_total_itens integer,
        valor_a_pagar decimal,
        valor_tributos decimal,
        numero varchar(50),
        serie varchar(50),
        emissao varchar(50),
        chave_de_acesso varchar(255),
        protocolo_autorizacao varchar(255),
        estabelecimento_id integer NOT NULL,
        consumidor_id integer,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_nota_fiscal_estabelecimento_id FOREIGN KEY (estabelecimento_id) REFERENCES estabelecimento (id),
        CONSTRAINT fk_nota_fiscal_consumidor_id FOREIGN KEY (consumidor_id) REFERENCES consumidor (id)
    );

CREATE TABLE
    unidade_medida (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

CREATE TABLE
    categoria_item (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

CREATE TABLE
    item (
        id int NOT NULL AUTO_INCREMENT,
        codigo varchar(60),
        descricao varchar(90) not null,
        unidade_medida_id integer not null,
        categoria_item_id integer,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_item_unidade_medida_id FOREIGN KEY (unidade_medida_id) REFERENCES unidade_medida (id),
        CONSTRAINT fk_item_categoria_item_id FOREIGN KEY (categoria_item_id) REFERENCES categoria_item (id)
    );

CREATE TABLE
    item_nota_fiscal (
        id int NOT NULL AUTO_INCREMENT,
        nota_fiscal_id integer not null,
        item_id integer not null,
        quantidade decimal not null,
        preco_unitario decimal not null,
        valor_total decimal not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_item_nota_fiscal_nota_fiscal_id FOREIGN KEY (nota_fiscal_id) REFERENCES nota_fiscal (id),
        CONSTRAINT fk_item_nota_fiscal_item_id FOREIGN KEY (item_id) REFERENCES item (id)
    );
////sss
CREATE TABLE
    item_estabelecimento (
        id int NOT NULL AUTO_INCREMENT,
        item_id integer not null,
        item_referencia_id integer,
        estabelecimento_id integer,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_item_estabelecimento_nota_fiscal_item_id FOREIGN KEY (item_id) REFERENCES item (id),
        CONSTRAINT fk_item_estabelecimento_item_referencia_id FOREIGN KEY (item_referencia_id) REFERENCES item (id),
        CONSTRAINT fk_item_estabelecimento_estabelecimento_id FOREIGN KEY (estabelecimento_id) REFERENCES estabelecimento (id)
    );

CREATE TABLE
    tipo_pagamento (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

CREATE TABLE
    forma_pagamento (
        id int NOT NULL AUTO_INCREMENT,
        codigo varchar(60),
        descricao varchar(90) not null,
        tipo_pagamento_id integer not null,
        nota_fiscal_id integer not null,
        valor_pago decimal not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_forma_pagamento_tipo_pagamento_id FOREIGN KEY (tipo_pagamento_id) REFERENCES tipo_pagamento (id),
        CONSTRAINT fk_forma_pagamento_nota_fiscal_id FOREIGN KEY (nota_fiscal_id) REFERENCES nota_fiscal (id)
    );

select
    e.descricao nome_fantasia,
    i.descricao as item_descricao,
    um.descricao as unidade_medida_descricao,
    inf.preco_unitario,
    inf.quantidade,
    inf.valor_total,
    c.descricao nome_consumidor,
    nf.data_emissao
from
    nota_fiscal nf
    join estabelecimento e on e.id = nf.estabelecimento_id
    left join consumidor c on c.id = nf.consumidor_id
    join item_nota_fiscal inf on inf.nota_fiscal_id = nf.id
    join item i on i.id = inf.item_id
    join unidade_medida um on um.id = i.unidade_medida_id
WHERE
    nf.data_emissao BETWEEN DATE_FORMAT ('2025-06-15', '%Y-%m-01') AND LAST_DAY  ('2025-06-15')
LIMIT
    0, 200;