DROP TABLE IF EXISTS project.tb_52_weeks CASCADE;
CREATE TABLE IF NOT EXISTS project.tb_52_weeks (
    "code_id" INTEGER NOT NULL REFERENCES "meta"."code"(id),
    "price_type" INTEGER NOT NULL REFERENCES "meta"."config"(id),
    "row_type" BOOLEAN NOT NULL ,
    "unit_type" INTEGER NOT NULL,
    "unit" INTEGER NOT NULL,
    "np_dt" INTEGER,
    "np_val" numeric(20, 2),
    "op_dt" INTEGER,
    "op_val" numeric(20, 2),
    "p_percent" numeric(10, 2)
);
    COMMENT ON COLUMN "project"."tb_52_weeks"."code_id" IS '코드ID';
    COMMENT ON COLUMN "project"."tb_52_weeks"."price_type" IS '가격종류';
    COMMENT ON COLUMN "project"."tb_52_weeks"."row_type" IS '1:최고가,0:최저가';
    COMMENT ON COLUMN "project"."tb_52_weeks"."unit_type" IS '기간 단위 1:주, 2:월';
    COMMENT ON COLUMN "project"."tb_52_weeks"."unit" IS '기간 값';
    COMMENT ON COLUMN "project"."tb_52_weeks"."np_dt" IS '현재 가격 일자';
    COMMENT ON COLUMN "project"."tb_52_weeks"."np_val" IS '현재 가격 값';
    COMMENT ON COLUMN "project"."tb_52_weeks"."op_dt" IS '과거 가격 일자';
    COMMENT ON COLUMN "project"."tb_52_weeks"."op_val" IS '과거 가격 값';
    COMMENT ON COLUMN "project"."tb_52_weeks"."p_percent" IS '과거와 현재 가격 차이 퍼센트';
    ----------------------------------------------