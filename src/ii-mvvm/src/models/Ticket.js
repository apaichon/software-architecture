export class Ticket {
  constructor(concertId, studentName, studentClass) {
    this.concertId = concertId;
    this.studentName = studentName;
    this.studentClass = studentClass;
    this.purchaseDate = new Date();
  }
}