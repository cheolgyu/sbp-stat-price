DROP TABLE IF EXISTS project.tb_52_weeks CASCADE;
CREATE TABLE IF NOT EXISTS project.tb_52_weeks (
    "code_id" INTEGER NOT NULL REFERENCES "meta"."code"(id),
    "price_type" INTEGER NOT NULL REFERENCES "meta"."config"(id),
    "unit_type" INTEGER NOT NULL,
    "unit" INTEGER NOT NULL,
    "high_price" numeric(20, 2),
    "low_price" numeric(20, 2)
);
