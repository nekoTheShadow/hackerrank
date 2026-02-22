#lang racket

(require "./main.rkt"
         rackunit)

(define (test-main path expected)
  (define (test-main-sub in)
    (define out (open-output-string))
    (parameterize ([current-input-port in]
                   [current-output-port out])
      (main)
      (check-equal? (get-output-string out) expected)))

  (call-with-input-file path test-main-sub))

(module+ test
  (test-main "example1.txt" "4 2 5 1 3\n"))
