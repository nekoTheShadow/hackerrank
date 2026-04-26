#lang racket

(define (main)
  (define nk (map string->number (string-split (read-line))))
  (define prices (map string->number (string-split (read-line))))
  (define n (list-ref nk 0))
  (define k (list-ref nk 1))

  (define count 0)
  (define total 0)
  (define answer
    (let/ec return
      (for ([price (sort prices <)])
        (when (< k (+ total price))
          (return count))
        (set! count (+ count 1))
        (set! total (+ total price)))
      (return count)))
  (displayln answer))
(main)

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
  (test-main "./example.txt" "4\n"))
