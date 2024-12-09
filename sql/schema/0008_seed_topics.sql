-- +goose Up
INSERT INTO topics (id, created_at, updated_at, name) VALUES 
    (gen_random_uuid(), NOW(), NOW(), 'Artificial Intelligence & Machine Learning'),
    (gen_random_uuid(), NOW(), NOW(), 'Blockchain & Cryptocurrency'),
    (gen_random_uuid(), NOW(), NOW(), 'Virtual & Augmented Reality'),
    (gen_random_uuid(), NOW(), NOW(), 'Internet of Things'),
    (gen_random_uuid(), NOW(), NOW(), 'Cloud Computing'),
    (gen_random_uuid(), NOW(), NOW(), 'Cybersecurity'),
    (gen_random_uuid(), NOW(), NOW(), 'Robotics'),
    (gen_random_uuid(), NOW(), NOW(), 'Big Data'),
    (gen_random_uuid(), NOW(), NOW(), 'Software Development'),
    (gen_random_uuid(), NOW(), NOW(), 'Quantum Computing'),
    (gen_random_uuid(), NOW(), NOW(), '3D Printing'),
    (gen_random_uuid(), NOW(), NOW(), 'Autonomous Vehicles'),
    (gen_random_uuid(), NOW(), NOW(), 'Renewable Energy'),
    (gen_random_uuid(), NOW(), NOW(), 'Nanotech'),
    (gen_random_uuid(), NOW(), NOW(), 'Biotech & Genomics'),
    (gen_random_uuid(), NOW(), NOW(), 'Solar Energy'),
    (gen_random_uuid(), NOW(), NOW(), 'Nuclear Energy'),
    (gen_random_uuid(), NOW(), NOW(), 'Space'),
    (gen_random_uuid(), NOW(), NOW(), 'Biology'),
    (gen_random_uuid(), NOW(), NOW(), 'Chemistry'),
    (gen_random_uuid(), NOW(), NOW(), 'Physics'),
    (gen_random_uuid(), NOW(), NOW(), 'Environmental Science'),
    (gen_random_uuid(), NOW(), NOW(), 'Earth Science'),
    (gen_random_uuid(), NOW(), NOW(), 'Marine Biology'),
    (gen_random_uuid(), NOW(), NOW(), 'Entrepreneurship'),
    (gen_random_uuid(), NOW(), NOW(), 'Startups'),
    (gen_random_uuid(), NOW(), NOW(), 'Marketing'),
    (gen_random_uuid(), NOW(), NOW(), 'E-commerce'),
    (gen_random_uuid(), NOW(), NOW(), 'Retail & Consumer Goods'),
    (gen_random_uuid(), NOW(), NOW(), 'Business Intelligence'),
    (gen_random_uuid(), NOW(), NOW(), 'Leadership'),
    (gen_random_uuid(), NOW(), NOW(), 'Sales'),
    (gen_random_uuid(), NOW(), NOW(), 'Mergers & Acquisitions'),
    (gen_random_uuid(), NOW(), NOW(), 'Human Resources'),
    (gen_random_uuid(), NOW(), NOW(), 'Residential Real Estate'),
    (gen_random_uuid(), NOW(), NOW(), 'Real Estate'),
    (gen_random_uuid(), NOW(), NOW(), 'Supply Chain Management'),
    (gen_random_uuid(), NOW(), NOW(), 'International Business'),
    (gen_random_uuid(), NOW(), NOW(), 'Product Management'),
    (gen_random_uuid(), NOW(), NOW(), 'Mental Health'),
    (gen_random_uuid(), NOW(), NOW(), 'Physical Health'),
    (gen_random_uuid(), NOW(), NOW(), 'Healthcare Technology'),
    (gen_random_uuid(), NOW(), NOW(), 'Aging & Longevity'),
    (gen_random_uuid(), NOW(), NOW(), 'Women Health'),
    (gen_random_uuid(), NOW(), NOW(), 'Men Health'),
    (gen_random_uuid(), NOW(), NOW(), 'Pediatrics'),
    (gen_random_uuid(), NOW(), NOW(), 'eSports'),
    (gen_random_uuid(), NOW(), NOW(), 'Sport'),
    (gen_random_uuid(), NOW(), NOW(), 'Elections'),
    (gen_random_uuid(), NOW(), NOW(), 'Political Theory'),
    (gen_random_uuid(), NOW(), NOW(), 'Immigration'),
    (gen_random_uuid(), NOW(), NOW(), 'Stock Market'),
    (gen_random_uuid(), NOW(), NOW(), 'Investing & Trading'),
    (gen_random_uuid(), NOW(), NOW(), 'Economics'),
    (gen_random_uuid(), NOW(), NOW(), 'Commercial Real Estate'),
    (gen_random_uuid(), NOW(), NOW(), 'Financial Technology'),
    (gen_random_uuid(), NOW(), NOW(), 'Movies & Film'),
    (gen_random_uuid(), NOW(), NOW(), 'Music'),
    (gen_random_uuid(), NOW(), NOW(), 'Video Games'),
    (gen_random_uuid(), NOW(), NOW(), 'Books & Literature'),
    (gen_random_uuid(), NOW(), NOW(), 'Travel & Adventure'),
    (gen_random_uuid(), NOW(), NOW(), 'Food & Cuisine'),
    (gen_random_uuid(), NOW(), NOW(), 'Fashion & Style'),
    (gen_random_uuid(), NOW(), NOW(), 'Photography & Videography'),
    (gen_random_uuid(), NOW(), NOW(), 'World Wars'),
    (gen_random_uuid(), NOW(), NOW(), 'Criminal Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Civil Law'),
    (gen_random_uuid(), NOW(), NOW(), 'International Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Legal Technology'),
    (gen_random_uuid(), NOW(), NOW(), 'Intellectual Property Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Family Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Legal Research & Writing'),
    (gen_random_uuid(), NOW(), NOW(), 'Environmental Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Human Rights Law'),
    (gen_random_uuid(), NOW(), NOW(), 'Climate Change'),
    (gen_random_uuid(), NOW(), NOW(), 'Sustainable Agriculture'),
    (gen_random_uuid(), NOW(), NOW(), 'Green Tech'),
    (gen_random_uuid(), NOW(), NOW(), 'Worlds Economics'),
    (gen_random_uuid(), NOW(), NOW(), 'America Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'America Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'China Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'China Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'India Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'India Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'European Union Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'European Union Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'Russia Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'Russia Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'Brazil Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'Brazil Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'Africa Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'Africa Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'Middle East Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'Middle East Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'Asia Politics'),
    (gen_random_uuid(), NOW(), NOW(), 'Asia Economy'),
    (gen_random_uuid(), NOW(), NOW(), 'War');

-- +goose Down
DELETE FROM topics;