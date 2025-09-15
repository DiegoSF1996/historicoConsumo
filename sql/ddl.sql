CREATE TABLE
    estabelecimento (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(50),
        endereco varchar(200),
        cnpj varchar(50),
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );
ALTER TABLE estabelecimento ADD UNIQUE KEY estabelecimento_descricao_cnpj (descricao,cnpj);

CREATE TABLE
    consumidor (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(150),
        cpf varchar(50),
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

ALTER TABLE consumidor ADD UNIQUE KEY consumidor_descricao_cpf (descricao,cpf);

CREATE TABLE
    nota_fiscal (
        id integer NOT NULL AUTO_INCREMENT,
        data_autorizacao datetime,
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

ALTER TABLE nota_fiscal ADD UNIQUE KEY nota_fiscal_numero_serie_emissao (numero,serie,emissao);

CREATE TABLE
    unidade_medida (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

ALTER TABLE unidade_medida ADD UNIQUE KEY unique_unidade_descricao (descricao);


CREATE TABLE
    categoria_item (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

ALTER TABLE categoria_item ADD UNIQUE KEY categoria_item_descricao (descricao);

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

ALTER TABLE item ADD UNIQUE KEY item_codigo_descricao_unidade (codigo,descricao,unidade_medida_id);

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

ALTER TABLE item_estabelecimento ADD UNIQUE KEY item_estabelecimento (item_id,estabelecimento_id);

CREATE TABLE
    tipo_pagamento (
        id int NOT NULL AUTO_INCREMENT,
        descricao varchar(90) not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID)
    );

    ALTER TABLE tipo_pagamento ADD UNIQUE KEY tipo_pagamento_descricao (descricao);

CREATE TABLE
    forma_pagamento_nota_fiscal (
        id int NOT NULL AUTO_INCREMENT,
        tipo_pagamento_id integer not null,
        nota_fiscal_id integer not null,
        valor_pago decimal not null,
        created_at datetime not null default current_timestamp,
        updated_at datetime not null default current_timestamp,
        PRIMARY KEY (ID),
        CONSTRAINT fk_forma_pagamento_tipo_pagamento_id FOREIGN KEY (tipo_pagamento_id) REFERENCES tipo_pagamento (id),
        CONSTRAINT fk_forma_pagamento_nota_fiscal_id FOREIGN KEY (nota_fiscal_id) REFERENCES nota_fiscal (id)
    );

//ALTER TABLE forma_pagamento_nota_fiscal ADD UNIQUE KEY forma_pagamento_nota_fiscal_nota_fiscal_tipo_pagamento (nota_fiscal_id,tipo_pagamento_id);

select
    e.descricao nome_fantasia,
    i.descricao as item_descricao,
    um.descricao as unidade_medida_descricao,
    inf.preco_unitario,
    inf.quantidade,
    inf.valor_total,
    c.descricao nome_consumidor,
    nf.data_autorizacao
from
    nota_fiscal nf
    join estabelecimento e on e.id = nf.estabelecimento_id
    left join consumidor c on c.id = nf.consumidor_id
    join item_nota_fiscal inf on inf.nota_fiscal_id = nf.id
    join item i on i.id = inf.item_id
    join unidade_medida um on um.id = i.unidade_medida_id
WHERE
    nf.data_autorizacao BETWEEN DATE_FORMAT ('2025-06-15', '%Y-%m-01') AND LAST_DAY  ('2025-06-15')
LIMIT
    0, 200;

// obtém o total gasto no ano corrente agrupado por estabelecimento
select
    coalesce(apelido,e.descricao) nome_fantasia, 
    sum(coalesce(nf.valor_tributos,0)) total_gasto_tributos,
    sum(coalesce(nf.valor_a_pagar,0)) total_gasto,
    sum(coalesce(nf.valor_tributos,0)) /  sum(coalesce(nf.valor_a_pagar,0)) * 100 percentual_imposto,
    data_autorizacao 
from nota_fiscal nf
join estabelecimento e on e.id = nf.estabelecimento_id
left join consumidor c on c.id = nf.consumidor_id
join item_nota_fiscal inf on inf.nota_fiscal_id = nf.id
join item i on i.id = inf.item_id
join unidade_medida um on um.id = i.unidade_medida_id
where  nf.data_autorizacao BETWEEN DATE_FORMAT (current_date , '%Y-01-01') AND  LAST_DAY(DATE_FORMAT(current_date, '%Y-12-01'))
group by e.descricao 
order by 3 desc 

// atualiza a data_autorizacao
update nota_fiscal 
set data_autorizacao = CASE
    WHEN instr(nf.data_autorizacao, '/') > 0
    THEN strftime(
      '%Y-%m-%d %H:%M:%S',
      substr(nf.data_autorizacao, 7, 4) || '-' ||   -- ano
      substr(nf.data_autorizacao, 4, 2) || '-' ||   -- mês
      substr(nf.data_autorizacao, 1, 2) || ' ' ||   -- dia
      substr(nf.data_autorizacao, 12, 8)            -- hora
    )
    ELSE nf.data_autorizacao
  END
from nota_fiscal nf 
where nf.id = nota_fiscal.id 