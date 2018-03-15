package stats

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type statsRepositorySQL struct{}

func (s *statsRepositorySQL) GetLeaderboardPlayersByLeagueUniqueName(transaction transaction.Transaction, uniqueName string) ([]LeaderboardPlayer, error) {
	leaderboardPlayers := []LeaderboardPlayer{}

	query := `
		SELECT
			p.display_name,
			r.rating,
			COUNT(CASE WHEN d.has_won THEN 1 END) AS won_match_count,
			COUNT(*) AS total_match_count
		FROM
			players p,
			ratings r,
			leagues l,
			matches m,
			details d
		WHERE
			l.unique_name = $1 AND
			l.id = r.league_id AND
			r.player_id = p.id AND
			p.id = d.player_id AND
			m.league_id = l.id AND
			d.match_id = m.id
		GROUP BY
			p.display_name,
			r.rating
		ORDER BY
			r.rating DESC;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	err := sqlTransaction.Select(&leaderboardPlayers, query, uniqueName)

	return leaderboardPlayers, err
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &statsRepositorySQL{}
}
