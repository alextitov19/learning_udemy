class QuizQuestion {
  const QuizQuestion(this.text, this.answers);

  final String text;
  final List<String> answers;

  List<String> getShuffledAnswers() {
    final l = List.of(answers);
    l.shuffle();
    return l;
  }
}
