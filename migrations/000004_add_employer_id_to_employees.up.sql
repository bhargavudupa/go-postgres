ALTER TABLE employees
ADD COLUMN employer_id INT,
ADD CONSTRAINT fk_employer
FOREIGN KEY (employer_id)
REFERENCES employers(id)
ON DELETE SET NULL;