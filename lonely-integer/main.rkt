#lang racket

(define (main)
  (define n (string->number (read-line)))
  (define a (map string->number (string-split (read-line))))

  (define ht (make-hash))
  (for ([v a])
    (hash-update! ht v add1 0))
  (for ([(k v) (in-hash ht)])
    (when (= v 1)
      (displayln k))))

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

  (test-main "./input0.txt" "1\n")
  (test-main "./input1.txt" "2\n")
  (test-main "./input2.txt" "2\n"))
