import java.util.*;

class Underwriter {
  
  public static String[] identifyInvalidTransactions(String[] transactions) {
      List<String> result = new ArrayList<>();
      Map<String, List<Transaction>> nameToTransactions = new HashMap<>();
      for (String transaction : transactions) {
          Transaction t = new Transaction(transaction);
          if (t.amount > 2000) {
              result.add(transaction);
          }
          List<Transaction> sameNameTransactions = nameToTransactions.getOrDefault(t.name, new ArrayList<>());
          for (Transaction prevTransaction : sameNameTransactions) {
              if (t.isInvalid(prevTransaction)) {
                  result.add(transaction);
                  result.add(prevTransaction.originalTransaction);
              }
          }
          sameNameTransactions.add(t);
          nameToTransactions.put(t.name, sameNameTransactions);
      }
      return result.toArray(new String[0]);
  }
    
    private static class Transaction {
        String originalTransaction;
        String name;
        int time;
        int amount;
        String city;

        public Transaction(String transaction) {
            this.originalTransaction = transaction;
            String[] parts = transaction.split(",");
            this.name = parts[0];
            this.time = Integer.parseInt(parts[1]);
            this.amount = Integer.parseInt(parts[2]);
            this.city = parts[3];
        }

        public boolean isInvalid(Transaction prevTransaction) {
            if (!this.name.equals(prevTransaction.name)) {
                return false;
            }
            if (Math.abs(this.time - prevTransaction.time) <= 60 && !this.city.equals(prevTransaction.city)) {
                return true;
            }
            if (Math.abs(this.time - prevTransaction.time) <= 60 && this.amount == prevTransaction.amount) {
                return true;
            }
            if (this.time == prevTransaction.time) {
                return true;
            }
            return false;
        }
    }
  
}