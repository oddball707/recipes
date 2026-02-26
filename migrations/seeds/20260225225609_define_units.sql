-- +goose Up
-- +goose StatementBegin

INSERT INTO units (name, abbreviation) VALUES ('gram', 'g');
INSERT INTO units (name, abbreviation) VALUES ('milliliter', 'ml');
INSERT INTO units (name, abbreviation) VALUES ('teaspoon', 'tsp');
INSERT INTO units (name, abbreviation) VALUES ('tablespoon', 'tbsp');
INSERT INTO units (name, abbreviation) VALUES ('cup', 'cup');
INSERT INTO units (name, abbreviation) VALUES ('ounce', 'oz');
INSERT INTO units (name, abbreviation) VALUES ('pound', 'lb');
INSERT INTO units (name, abbreviation) VALUES ('liter', 'l');
INSERT INTO units (name, abbreviation) VALUES ('milligram', 'mg');
INSERT INTO units (name, abbreviation) VALUES ('pinch', 'pinch');
INSERT INTO units (name, abbreviation) VALUES ('dash', 'dash');
INSERT INTO units (name, abbreviation) VALUES ('quart', 'qt');
INSERT INTO units (name, abbreviation) VALUES ('gallon', 'gal');
INSERT INTO units (name, abbreviation) VALUES ('pieces', 'pcs');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM units WHERE name IN (
  'gram',
  'milliliter',
  'teaspoon',
  'tablespoon',
  'cup',
  'ounce',
  'pound',
  'liter',
  'milligram',
  'pinch',
  'dash',
  'quart',
  'gallon'
);
-- +goose StatementEnd

