ALTER TABLE ad_campaign ADD column if not exists is_started boolean default false not null;
ALTER TABLE ad_campaign ADD column if not exists date_started DATE;

ALTER TABLE ad_campaign ADD column if not exists date_created DATE NOT NULL DEFAULT CURRENT_DATE;