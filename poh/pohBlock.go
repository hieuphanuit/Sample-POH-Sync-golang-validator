package poh

import pb "example_poh.com/proto"

func GetTotalTransaction(block *pb.POHBlock) int {
	totalTransaction := 0
	for _, tick := range block.Ticks {
		for _, hash := range tick.Hashes {
			totalTransaction += len(hash.Transactions)
		}
	}

	return totalTransaction
}

func GetTransactions(block *pb.POHBlock) []*pb.Transaction {
	transactions := []*pb.Transaction{}
	for _, tick := range block.Ticks {
		for _, hash := range tick.Hashes {
			transactions = append(transactions, hash.Transactions...)
		}
	}
	return transactions
}
