-- +goose Up
-- +goose StatementBegin
INSERT INTO ingredients_categories (name, description) 
VALUES 
  ('vegetable', 'A plant or part of a plant used as food'),
  ('fruit', 'The mature ovary of a flowering plant'),
  ('protein', 'A source of amino acids'),
  ('dairy', 'Food produced from the milk of mammals'),
  ('spice', 'A substance used to flavor food'),
  ('grain', 'A small, hard, dry seed, with or without an attached hull or fruit layer, harvested for human or animal consumption'),
  ('other', 'Other types of ingredients'),
  ('condiment', 'A substance, such as a sauce or seasoning, added to food to enhance its flavor'),
  ('legume', 'The fruit or seed of any of various bean or pea plants');


INSERT INTO ingredients (name, is_system_ingredient, user_id, category_id)
VALUES
  ('steak', true, 1, 3), -- protein
  ('chicken breast', true, 1, 3), -- protein
  ('chicken thigh', true, 1, 3), -- protein
  ('skinless chicken thigh', true, 1, 3), -- protein
  ('skinless boneless chicken thigh', true, 1, 3), -- protein
  ('white rice', true, 1, 6), -- grain
  ('black beans', true, 1, 9), -- legume
  ('carrot', true, 1, 1), -- vegetable
  ('onion', true, 1, 1), -- vegetable
  ('salt', true, 1, 5), -- spice
  ('black pepper', true, 1, 5), -- spice
  ('smoked paprika', true, 1, 5), -- spice
  ('sweet paprika', true, 1, 5), -- spice
  ('spicy paprika', true, 1, 5), -- spice
  ('white sugar', true, 1, 7), -- other
  ('orange peel', true, 1, 2), -- fruit
  ('lemon peel', true, 1, 2), -- fruit
  ('mango', true, 1, 2), -- fruit
  ('apple', true, 1, 2), -- fruit
  ('green apple', true, 1, 2), -- fruit
  ('blueberry', true, 1, 2), -- fruit
  ('blackberry', true, 1, 2), -- fruit
  ('raspberry', true, 1, 2), -- fruit
  ('strawberry', true, 1, 2), -- fruit
  ('banana', true, 1, 2), -- fruit
  ('garlic', true, 1, 1), -- vegetable
  ('bell pepper', true, 1, 1), -- vegetable
  ('spinach', true, 1, 1), -- vegetable
  ('zucchini', true, 1, 1), -- vegetable
  ('broccoli', true, 1, 1), -- vegetable
  ('cauliflower', true, 1, 1), -- vegetable
  ('mushroom', true, 1, 1), -- vegetable
  ('ginger', true, 1, 5), -- spice
  ('turmeric', true, 1, 5), -- spice
  ('chili powder', true, 1, 5), -- spice
  ('cumin', true, 1, 5), -- spice
  ('nutmeg', true, 1, 5), -- spice
  ('brown sugar', true, 1, 7), -- other
  ('honey', true, 1, 7), -- other
  ('maple syrup', true, 1, 7), -- other
  ('olive oil', true, 1, 8), -- condiment
  ('soy sauce', true, 1, 8), -- condiment
  ('tomato sauce', true, 1, 8), -- condiment
  ('mustard', true, 1, 8), -- condiment
  ('mayonnaise', true, 1, 8), -- condiment
  ('vinegar', true, 1, 8), -- condiment
  ('peach', true, 1, 2), -- fruit
  ('pear', true, 1, 2), -- fruit
  ('grape', true, 1, 2), -- fruit
  ('pineapple', true, 1, 2), -- fruit
  ('watermelon', true, 1, 2), -- fruit
  ('almond', true, 1, 7), -- other
  ('walnut', true, 1, 7), -- other
  ('peanut', true, 1, 7), -- other
  ('chia seed', true, 1, 7), -- other
  ('flax seed', true, 1, 7), -- other
  ('celery', true, 1, 1), -- vegetable
  ('parsley', true, 1, 1), -- vegetable
  ('basil', true, 1, 1), -- vegetable
  ('rosemary', true, 1, 5), -- spice
  ('cayenne pepper', true, 1, 5), -- spice
  ('tarragon', true, 1, 5), -- spice
  ('dill', true, 1, 5), -- spice
  ('capers', true, 1, 8), -- condiment
  ('thyme', true, 1, 5), -- spice
  ('red wine vinegar', true, 1, 8), -- condiment
  ('balsamic vinegar', true, 1, 8), -- condiment
  ('sesame oil', true, 1, 8), -- condiment
  ('worcestershire sauce', true, 1, 8), -- condiment
  ('barbecue sauce', true, 1, 8), -- condiment
  ('lentils', true, 1, 9), -- legume
  ('tofu', true, 1, 3), -- protein
  ('tempeh', true, 1, 3), -- protein
  ('coconut oil', true, 1, 8), -- condiment
  ('saffron', true, 1, 5), -- spice
  ('agave syrup', true, 1, 7), -- other
  ('coconut sugar', true, 1, 7), -- other
  ('cashews', true, 1, 7), -- other
  ('sesame seeds', true, 1, 7), -- other
  ('sunflower seeds', true, 1, 7), -- other
  ('pumpkin seeds', true, 1, 7), -- other
  ('avocado', true, 1, 1), -- vegetable
  ('sriracha', true, 1, 8), -- condiment
  ('hoisin sauce', true, 1, 8), -- condiment
  ('asparagus', true, 1, 1), -- vegetable
  ('eggplant', true, 1, 1); -- vegetable
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ingredients WHERE is_system_ingredient = true AND user_id = 1;
DELETE FROM ingredients_categories WHERE id BETWEEN 1 AND 9
-- +goose StatementEnd
