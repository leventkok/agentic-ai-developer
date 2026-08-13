/*
 DAY 10 — Before/After

 Before (public fields, no rules):
   - Anyone could set activeLoans to 99
   - Overdue member could still borrow
   - Return a book you never borrowed

 After (encapsulated):
   - Loan list only changes via borrow/returnBook
   - Max 3 books enforced
   - Overdue blocks new borrows
*/

class Book {
  final String title;
  final String isbn;

  Book(this.title, this.isbn){
    if(title.isEmpty) throw ArgumentError('Title cannot be empty');
    if(isbn.isEmpty) throw ArgumentError('ISBN cannot be empty');
  }
}


class Loan {
  final Book book;
  final DateTime dueDate;


  



  Loan(this.book, this.dueDate){
    if(dueDate.isBefore(DateTime.now())) throw ArgumentError('Due date cannot be in the past');
    if(book.isbn.isEmpty) throw ArgumentError('ISBN cannot be empty');

  }

  bool isOverdue() => DateTime.now().isAfter(dueDate);
  String describe() => 'Loan(book: ${book.title}, dueDate: $dueDate)';


  
  
}


class Member {
  final String _name;
  final List<Loan> _activeLoans = [];

  Member(this._name){
    if(_name.isEmpty) throw ArgumentError('Name cannot be empty');
  }

  String get name => _name;
  List<Loan> get activeLoans => List.unmodifiable(_activeLoans);

  void borrow(Book book) {
    if (_activeLoans.any((loan) => loan.isOverdue())) throw StateError('Has overdue book');
    if (_activeLoans.length >= 3) throw StateError('Max 3 books');
    final due = DateTime.now().add(Duration(days: 14));
    _activeLoans.add(Loan(book, due));
  }

  int get activeLoanCount => _activeLoans.length;


 void returnBook(Book book) {
  final index = _activeLoans.indexWhere((loan) => loan.book.isbn == book.isbn);
  if (index == -1) throw StateError('Book not borrowed');
  _activeLoans.removeAt(index);
}



}


void main() {
  final member = Member('Levent');
  final book = Book('Clean Code', '978-0132350884');
  member.borrow(book);
  print(member.activeLoanCount);  


  print('=== Max 3 books ===');

  print('=== Return book ===');
  member.returnBook(book);
  print(member.activeLoanCount);  

  print('=== Return not borrowed ===');
  try { member.returnBook(Book('X', '123')); } catch (e) { print('Error: $e'); }
}