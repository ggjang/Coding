state = {}
def apply(metric):
  # repeat for all recipes until no changes have been made
  # divide recipe into target, equation
  # parse equation into calculator with following rules.
  # 1. reserved keywords : INTERVAL
  # 2. strings are considered 'field keys' to fetch 'field values'
  # 3. only simple arithmatic is performed
  modified = True
  while modified:
    modified = False
    for recipe in recipes:
      recipe_toks = recipe.replace(' ','').split("=")
      if len(recipe_toks) != 2: fail("Invalid recipe format")
      target, equation = recipe_toks
      if target == "" or equation == "": fail("Invalid recipe format")
      # if target is already in metric, we can assume this recipe was performed. PASS
      if target in metric.fields.keys():
        continue
      # check if tokens are all available
      components = getComponents(equation)
      if not isSolvable(components, metric):
        continue
      # target is not in metric. We first need to check if tokens are available
      toks = tokenize(equation)
      toks = toPostFix(toks)
      res = evalPostFix(toks, metric)
      metric.fields[target] = res
      modified = True
  # drop any _diff fields from metric
  tmp = deepcopy(metric)
  for field in tmp.fields:
    if "_diff" in field:
      metric.fields.pop(field,None)
    
  return metric

def evalPostFix(toks, metric):
  ops = ['+','-','*','/']
  stack = []

  for tok in toks:
    if tok in ops:
      b = str2val(stack.pop(), metric)
      a = str2val(stack.pop(), metric)
      if tok == '+':
        res = a+b
      elif tok == '-':
        res = a-b
      elif tok == '*':
        res = a*b
      elif tok == '/':
        res = a/b
      else:
	fail("Unknown operator")
      stack.append(res)
    else:
      stack.append(tok)

  if len(stack) != 1: fail("Invalid recipe format")
  return stack[-1]

def str2val(s, metric): # string to value
  if type(s) != "string": return s
  digit = '0123456789'
  lower = 'abcdefghijklmnopqrstuvwxyz'
  upper = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  if s[0] in digit: # int
    return int(s)
  elif s[0] in lower: # field value
    return metric.fields[s]
  elif s[0] in upper: # KEYWORD
    if s == 'INTERVAL':
      if INTERVAL == 0: fail("INTERVAL cannot be 0")
      return INTERVAL
    else:
      fail("Unknown keyword referenced")
      return

def toPostFix(toks):
  parens = ['(',')']
  ops = ['+','-','*','/']
  letter = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
  digit = '0123456789'

  stack = []
  postfix = []

  for tok in toks:
    if tok not in ops+parens:
      postfix.append(tok)
    elif tok in ops:
      while len(stack) > 0 and priority(stack[-1]) >= priority(tok):
        postfix.append(stack.pop())
      stack.append(tok)
    elif tok == '(':
      stack.append(tok)
    elif tok == ')':
      while len(stack) > 0 and stack[-1] != '(':
        postfix.append(stack.pop())
      stack.pop()
  while len(stack) > 0:
    postfix.append(stack.pop())
  return postfix

def priority(op):
  if op == '+' or op == '-':
    return 1
  elif op == '*' or op == '/':
    return 2
  else:
    return 0


# is solvable IF
# components in equation are available    
def isSolvable(components, metric):
  for component in components:
    if component not in metric.fields.keys(): return False
  return True
    
def getComponents(equation):
  lower = 'abcdefghijklmnopqrstuvwxyz'
  tokens = tokenize(equation)
  components = []
  for token in tokens:
    if token[0] in lower: components.append(token)
  return components

def tokenize(equation):
  ops = ['+','-','*','/','(',')']
  letter = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
  digit = '0123456789'

  toks = []
  cur = ''

  for i in range(len(equation)):
    c = equation[i]
    if c == ' ': continue # ignore empty spaces

    if c in ops: # save previous if any
      if len(cur) > 0: toks.append(cur)
      toks.append(c)
      cur = ''
    else:
      cur += c
  if len(cur) > 0: toks.append(cur)
  return toks
