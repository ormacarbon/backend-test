CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    points INTEGER DEFAULT 0,
    share_code VARCHAR(50) UNIQUE,
    referred_by VARCHAR(50)
);

CREATE INDEX IF NOT EXISTS idx_users_share_code ON users(share_code);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_points ON users(points DESC);

INSERT INTO users (name, email, phone, points, share_code, referred_by) VALUES
    ('John Smith', 'john.smith@example.com', '+554512345678', 15, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NULL),
    ('Maria Garcia', 'maria.garcia@example.com', '+554519876543', 25, 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'),
    ('Alex Wong', 'alex.wong@example.com', '+554511223344', 10, 'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12'),
    ('Sarah Johnson', 'sarah.j@example.com', '+554514455667', 30, 'd3eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', NULL),
    ('Lucas Silva', 'lucas.silva@example.com', '+554519988776', 20, 'e4eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'd3eebc99-9c0b-4ef8-bb6d-6bb9bd380a14'),
    ('Emma Brown', 'emma.b@example.com', '+554513334444', 35, 'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', NULL),
    ('Carlos Rodriguez', 'carlos.r@example.com', '+554515556666', 40, 'g6eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16'),
    ('Ana Santos', 'ana.s@example.com', '+554517778888', 22, 'h7eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', NULL),
    ('David Kim', 'david.k@example.com', '+554512223333', 18, 'i8eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'h7eebc99-9c0b-4ef8-bb6d-6bb9bd380a18'),
    ('Isabella Chen', 'isabella.c@example.com', '+554514445555', 28, 'j9eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', NULL),
    ('Thomas Lee', 'thomas.l@example.com', '+554516667777', 33, 'k10eebc99-9c0b-4ef8-bb6d-6bb9bd380a21', 'j9eebc99-9c0b-4ef8-bb6d-6bb9bd380a20'),
    ('Julia Costa', 'julia.c@example.com', '+554518889999', 45, 'l11eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', NULL),
    ('Miguel Martinez', 'miguel.m@example.com', '+554511112222', 27, 'm12eebc99-9c0b-4ef8-bb6d-6bb9bd380a23', 'l11eebc99-9c0b-4ef8-bb6d-6bb9bd380a22'),
    ('Sophie Wilson', 'sophie.w@example.com', '+554513334444', 38, 'n13eebc99-9c0b-4ef8-bb6d-6bb9bd380a24', NULL),
    ('Daniel Park', 'daniel.p@example.com', '+554515556666', 42, 'o14eebc99-9c0b-4ef8-bb6d-6bb9bd380a25', 'n13eebc99-9c0b-4ef8-bb6d-6bb9bd380a24');