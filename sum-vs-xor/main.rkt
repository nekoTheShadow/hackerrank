#lang racket

(define (main)
  (define n (string->number (read-line)))
  (if (= n 0)
      (displayln 1)
      (displayln (expt 2 (count-0 n)))))

(define (count-0 x)
  (cond
    [[= x 0] 1]
    [[= x 1] 0]
    [else (+ (if (= (modulo x 2) 0) 1 0) (count-0 (quotient x 2)))]))

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

  (test-main "./input0.txt" "2\n")
  (test-main "./input1.txt" "4\n"))
