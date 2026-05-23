#lang racket

(define (main)
  (define q (string->number (read-line)))
  (define ss
    (for/list ([_ (in-range q)])
      (read-line)))
  (for ([s ss])
    (define tt (regexp-match* #rx"A+|B+" s))
    (displayln (for/sum ([t tt]) (- (string-length t) 1)))))

(module+ main
  (main))

(module+ test
  (require rackunit)
  (define (test-main filename expected)
    (define (run)
      (call-with-output-string run-with-out))

    (define (run-with-out stdout)
      (call-with-input-file filename (lambda (stdin) (run-with-in-out stdin stdout))))

    (define (run-with-in-out stdin stdout)
      (parameterize ([current-input-port stdin]
                     [current-output-port stdout])
        (main)))

    (check-equal? (run) expected))

  (test-main "./input.txt" (file->string "./output.txt")))
