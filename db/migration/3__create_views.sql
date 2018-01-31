CREATE OR REPLACE VIEW aggregated_venue_score AS
SELECT
    v.id as "venue_id",
    AVG(rate.total_score) as "score"
FROM ratings as rate
INNER JOIN venues AS v
ON rate.venue_id = v.id
GROUP by v.id;
