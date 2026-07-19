#lang racket

(define (main)
  (define T1T2N (map string->number (string-split (read-line))))
  (define T1 (list-ref T1T2N 0))
  (define T2 (list-ref T1T2N 1))
  (define N (list-ref T1T2N 2))

  (for ([_i (in-range (- N 1))])
    (define T3 (+ T1 (* T2 T2)))
    (set! T1 T2)
    (set! T2 T3))

  (displayln T1))

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
  (test-main "./00_input.txt" "5\n"))
