#lang racket

(define (main)
  (define t (string->number (read-line)))
  (define nn
    (for/list ([i (in-range t)])
      (string->number (read-line))))
  (for ([n nn])
    (displayln (solve n "Louise" "Richard"))))

(define (solve n a b)
  (let/ec return
    (when (= n 1)
      (return b))

    (define m (expt 2 (len n)))
    (if (= n m)
        (solve (quotient n 2) b a)
        (solve (- n m) b a))))

(define (len x)
  (if (= x 1)
      0
      (+ 1 (len (quotient x 2)))))

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

  (test-main "./input.txt" "Richard\n"))
