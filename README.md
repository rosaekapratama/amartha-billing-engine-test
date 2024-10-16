# amartha-billing-engine-test

Initialize database:
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.billings (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	loan_id uuid NOT NULL,
	installment_no int NOT NULL,
	installment_principal_amount numeric(15, 3) NOT NULL,
	installment_payment_amount numeric(15, 3) NULL,
	installment_interest_amount numeric(15, 3) NULL,
	interest_rate numeric(7, 5) NOT NULL,
	due_dt timestamptz NOT NULL,
	CONSTRAINT loans_pk PRIMARY KEY (id),
	CONSTRAINT loan_id_installment_no_un UNIQUE (loan_id, installment_no)
);

CREATE TABLE public.payments (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	loan_id uuid NOT NULL,
	payment_amount numeric(15, 3) NOT NULL,
	payment_dt timestamptz NOT NULL,
	CONSTRAINT payments_pk PRIMARY KEY (id)
);

INSERT INTO public.billings (id, loan_id, installment_no, interest_rate, installment_principal_amount, installment_payment_amount, installment_interest_amount, due_dt) VALUES
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 1, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '1 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')), -- Next Monday at EOD
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 2, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '2 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 3, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '3 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 4, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '4 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 5, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '5 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 6, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '6 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 7, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '7 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 8, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '8 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 9, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '9 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 10, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '10 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 11, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '11 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 12, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '12 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 13, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '13 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 14, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '14 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 15, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '15 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 16, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '16 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 17, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '17 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 18, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '18 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 19, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '19 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 20, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '20 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 21, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '21 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')), 
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 22, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '22 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 23, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '23 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 24, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '24 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 25, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '25 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 26, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '26 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 27, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '27 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 28, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '28 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 29, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '29 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 30, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '30 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 31, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '31 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 32, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '32 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 33, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '33 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 34, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '34 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 35, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '35 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 36, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '36 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 37, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '37 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 38, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '38 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 39, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '39 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 40, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '40 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 41, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '41 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 42, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '42 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 43, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '43 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 44, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '44 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 45, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '45 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 46, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '46 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 47, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '47 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 48, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '48 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 49, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '49 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds')),
    (uuid_generate_v4(), 'd3295395-468d-422f-b2dd-ccd7328c3767', 50, 0.1, 100000, 110000, 10000, (DATE_TRUNC('week', NOW()) + INTERVAL '50 week' + INTERVAL '23 hours' + INTERVAL '59 minutes' + INTERVAL '59 seconds'));
```

Adjust database name in **conf/app.yaml**
```
database:
  playground:
    database: playground <<< adjust accordingly
```

Get Outstanding Amount by Loan ID API
```
curl --location 'http://localhost:9092/v1/billing/d3295395-468d-422f-b2dd-ccd7328c3767/outstanding' \
--header 'Content-Type: application/json'
```

Check Loan Delinquent Status by Loan ID API
```
curl --location 'http://localhost:9092/v1/billing/d3295395-468d-422f-b2dd-ccd7328c3767/status' \
--header 'Content-Type: application/json'
```

Make Payment API
```
curl --location 'http://localhost:9092/v1/billing/payment' \
--header 'Content-Type: application/json' \
--data '{
    "loanId": "d3295395-468d-422f-b2dd-ccd7328c3767",
    "paymentAmount": 110000,
    "paymentDt": "2024-10-27T04:00:00+07:00"
}'
```
