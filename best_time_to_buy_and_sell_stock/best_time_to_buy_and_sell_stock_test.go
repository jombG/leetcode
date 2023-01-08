package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBestTime_1(t *testing.T) {
	require.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestBestTime_2(t *testing.T) {
	require.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestBestTime_3(t *testing.T) {
	require.Equal(t, 9, maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
}
