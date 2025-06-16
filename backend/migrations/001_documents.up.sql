-- Create Table of documents, rule groups and rules

CREATE TABLE Documents (
       id UUID PRIMARY KEY,
       payload JSON,
       combine_with_or BOOLEAN NOT NULL,
       create_at TIMESTAMP NOT NULL DEFAULT now(),
       update_at TIMESTAMP NOT NULL DEFAULT now(),
)


CREATE TABLE Groups (
       id UUID PRIMARY KEY,
       focument_id UUID REFERENCES Documents(Id) ON DELETE CASCADE,
       combine_with_or BOOLEAN NOT NULL
)

CREATE TABLE Rules (
       id UUID PRIMARY KEY,
       group_id UUID REFERENCES Groups(Id) ON DELETE CASCADE,
       field TEXT NOT NULL, -- since its a text field can contain custome field name as well
       operator TEXT NOT NULL,
       value_list TEXT[] NOT NUll
)
